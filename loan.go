package main

import (
	"log"
	"time"
)

type LoanInterest struct {
	Day             time.Time
	Balance         float64
	InterestRate    float64
	DaysInterest    float64
	RunningInterest float64
}

func DailyInterest(from, until time.Time) ([]LoanInterest, error) {
	loanAccountId := 2
	offsetAccountId := 1

	query := `SELECT principal.day, principal.balance, rates.interest_rate, round(principal.balance * ( rates.interest_rate / 100 / 365 ), 2) as days_interest,
    SUM(round(principal.balance * ( rates.interest_rate /100 / 365 ), 2)) OVER (ORDER BY principal.day)
    FROM
        (
        SELECT loan.day::date as day, SUM(loan.balance + savings.balance) as balance
        FROM
        (
            SELECT day::date,
            (SELECT starting_balance FROM accounts WHERE id = $1) + sum(COALESCE(day_balance, 0)) over (ORDER BY day::date) as balance
            FROM
            (
                SELECT day::date, sum(debit) AS day_debit, sum(credit) AS day_credit, sum(credit - debit) AS day_balance
                FROM generate_series($3::date, $4::date, '1 day') sq (day)
                LEFT JOIN transactions t
                ON day::date = occurred AND t.account_id = $1
                GROUP BY day::date
                ORDER BY day::date
            ) AS data1
        ) as loan
        INNER JOIN
        (
            SELECT day::date,
            (SELECT starting_balance FROM accounts WHERE id = $2) + sum(COALESCE(day_balance, 0)) over (ORDER BY day::date) as balance
            FROM
            (
                SELECT day::date, sum(debit) AS day_debit, sum(credit) AS day_credit, sum(credit - debit) AS day_balance
                FROM generate_series($3::date, $4::date, '1 day') sq (day)
                LEFT JOIN transactions t
                ON day::date = occurred AND t.account_id = $2
                GROUP BY day::date
                ORDER BY day::date
            ) AS data2
        ) as savings
        ON (savings.day::date = loan.day::date)

        GROUP BY loan.day::date
    ) as principal
    INNER JOIN
    (
        SELECT day::date,
        (SELECT rate FROM interest_rates WHERE effective_date <= day::date ORDER BY effective_date LIMIT 1) as interest_rate
        FROM generate_series($3::date, $4::date, '1 day') sq (day)
    ) as rates
    ON (rates.day::date = principal.day::date)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var day time.Time
	var balance, interest_rate, days_interest, running_interest float64

	var returnLines []LoanInterest

	rows, err := stmt.Query(loanAccountId, offsetAccountId, from, until)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&day,
			&balance,
			&interest_rate,
			&days_interest,
			&running_interest,
		)
		if err != nil {
			log.Printf("%s\n", err)
			return nil, err
		}

		// log.Printf(
		// 	"On %s balance was %.2f at a rate of %.2f, interest for the day was %.2f for a total to date of %2.f\n",
		// 	day,
		// 	balance,
		// 	interest_rate,
		// 	days_interest,
		// 	running_interest,
		// )

		l := &LoanInterest{
			Day:             day,
			Balance:         balance,
			InterestRate:    interest_rate,
			DaysInterest:    days_interest,
			RunningInterest: running_interest,
		}

		returnLines = append(returnLines, *l)
	}
	return returnLines, nil
}
