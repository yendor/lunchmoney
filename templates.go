package main

// Generated by templeGen. DO NOT EDIT!
var templates = map[string]string{
	"accounts_index.html": "e3t0ZW1wbGF0ZSAiaGVhZGVyLmh0bWwiIC59fQoKe3sgd2l0aCAkYWNjIDo9IC5BY2NvdW50IH19Cgo8aDM+e3sgJGFjYy5OYW1lIH19PC9oMz4KCjxkaXYgY2xhc3M9InRvdGFscyByb3ciPgogICAgPGRpdiBjbGFzcz0idG90YWwgY29sLW1kLTIiPgogICAgICAgIFRvdGFsOiB7eyAkYWNjLkdldEZvcm1hdHRlZFRvdGFsIH19IHt7ICRhY2MuQ3VycmVuY3lDb2RlIH19CiAgICA8L2Rpdj4KICAgIDxkaXYgY2xhc3M9InRvdGFsLWNsZWFyZWQgY29sLW1kLTMiPgogICAgICAgIENsZWFyZWQgVG90YWw6ICB7eyAkYWNjLkdldEZvcm1hdHRlZENsZWFyZWRUb3RhbCB9fSB7eyAkYWNjLkN1cnJlbmN5Q29kZSB9fQogICAgPC9kaXY+CjwvZGl2PgoKPGRpdiBjbGFzcz0iYWN0aW9ucyI+CiAgICA8Zm9ybSBhY3Rpb249Ii9hY2NvdW50cy9pbXBvcnQiIG1ldGhvZD0iUE9TVCIgZW5jdHlwZT0ibXVsdGlwYXJ0L2Zvcm0tZGF0YSI+CiAgICAgICAgPGRpdiBjbGFzcz0icm93Ij4KICAgICAgICAgICAgPGRpdiBjbGFzcz0iY29sLW1kLTYiPgogICAgICAgICAgICAgICAgPGlucHV0IHR5cGU9ImhpZGRlbiIgbmFtZT0iYWNjb3VudF9pZCIgdmFsdWU9Int7ICRhY2MuSUQgfX0iIC8+CiAgICAgICAgICAgICAgICA8aW5wdXQgdHlwZT0iZmlsZSIgbmFtZT0idXBsb2FkX2ZpbGUiIGNsYXNzPSJmb3JtLWNvbnRyb2wiIC8+CiAgICAgICAgICAgIDwvZGl2PgogICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb2wtbWQtNiI+CiAgICAgICAgICAgICAgICA8YnV0dG9uIHR5cGU9InN1Ym1pdCIgY2xhc3M9ImJ0biBidG4tZGVmYXVsdCI+SW1wb3J0PC9idXR0b24+CiAgICAgICAgICAgIDwvZGl2PgogICAgICAgIDwvZGl2PgogICAgPC9mb3JtPgo8L2Rpdj4KCjxkaXYgY2xhc3M9InJvdyI+CiAgICA8ZGl2IGNsYXNzPSJjb2wtbWQtMTIiPgogICAgICAgIDx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtc3RyaXBlZCB0YWJsZS1ib3JkZXJlZCBlZGl0YWJsZS10YWJsZSI+CiAgICAgICAgPHRoZWFkPgogICAgICAgIDx0cj4KICAgICAgICAgICAgPHRoPklEPC90aD4KICAgICAgICAgICAgPHRoPkRhdGU8L3RoPgogICAgICAgICAgICA8dGg+UGF5ZWU8L3RoPgogICAgICAgICAgICA8dGg+TWVtbzwvdGg+CiAgICAgICAgICAgIDx0aD5EZWJpdDwvdGg+CiAgICAgICAgICAgIDx0aD5DcmVkaXQ8L3RoPgogICAgICAgIDwvdHI+CiAgICAgICAgPC90aGVhZD4KICAgICAgICA8dGZvb3Q+CiAgICAgICAgPHRyPgogICAgICAgICAgICA8dGQ+PC90ZD4KICAgICAgICAgICAgPHRkPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48L3RkPgogICAgICAgICAgICA8dGQ+PC90ZD4KICAgICAgICAgICAgPHRkPjwvdGQ+CiAgICAgICAgPC90cj4KICAgICAgICA8L3Rmb290PgogICAgICAgIDx0Ym9keT4KICAgICAgICB7eyByYW5nZSAkYWNjLlRyYW5zYWN0aW9ucyB9fQogICAgICAgICAgICA8dHI+CiAgICAgICAgICAgICAgICA8dGQ+e3sgLklEIH19PC90ZD4KICAgICAgICAgICAgICAgIDx0ZD57eyAuRGF0ZSB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQ+e3sgLlBheWVlIH19PC90ZD4KICAgICAgICAgICAgICAgIDx0ZD57eyAuTWVtbyB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQgY2xhc3M9InRleHQtcmlnaHQiPnt7ICRhY2MuR2V0Rm9ybWF0dGVkQW1vdW50IC5EZWJpdCB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQgY2xhc3M9InRleHQtcmlnaHQiPnt7ICRhY2MuR2V0Rm9ybWF0dGVkQW1vdW50IC5DcmVkaXQgfX08L3RkPgogICAgICAgICAgICA8L3RyPgogICAgICAgIHt7IGVuZCB9fQogICAgICAgIDwvdGJvZHk+CiAgICAgICAgPC90YWJsZT4KICAgIDwvZGl2Pgo8L2Rpdj4KCjxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0IiBzcmM9Ii9wdWJsaWMvcGx1Z2lucy9qcXVlcnktdGFibGVkaXQvanF1ZXJ5LnRhYmxlZGl0Lm1pbi5qcyI+PC9zY3JpcHQ+Cgo8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+CiAgICAkKGRvY3VtZW50KS5yZWFkeShmdW5jdGlvbigpIHsKICAgICAgICAkKCcuZWRpdGFibGUtdGFibGUnKS5UYWJsZWRpdCh7CiAgICAgICAgICAgIHVybDogJy9hY2NvdW50cy97eyAkYWNjLklEIH19JywKICAgICAgICAgICAgaGlkZUlkZW50aWZpZXI6IHRydWUsCiAgICAgICAgICAgIGVkaXRCdXR0b246IHRydWUsCiAgICAgICAgICAgIHJlc3RvcmVCdXR0b246IGZhbHNlLAogICAgICAgICAgICBkZWxldGVCdXR0b246IGZhbHNlLAogICAgICAgICAgICBjb2x1bW5zOiB7CiAgICAgICAgICAgICAgICBpZGVudGlmaWVyOiBbMCwgJ0lEJ10sCiAgICAgICAgICAgICAgICBlZGl0YWJsZTogW1sxLCAnUGF5ZWUnXSwgWzIsICdNZW1vJ10sIFszLCAnRGViaXQnXSwgWzQsICdDcmVkaXQnXV0KICAgICAgICAgICAgfSwKICAgICAgICAgICAgYnV0dG9uczogewogICAgICAgICAgICAgICAgZWRpdDogewogICAgICAgICAgICAgICAgICAgIGNsYXNzOiAnYnRuIGJ0bi1zbSBidG4tcHJpbWFyeScsCiAgICAgICAgICAgICAgICAgICAgaHRtbDogJzxzcGFuIGNsYXNzPSJnbHlwaGljb24gZ2x5cGhpY29uLXBlbmNpbCI+PC9zcGFuPiAmbmJzcCBFRElUJywKICAgICAgICAgICAgICAgICAgICBhY3Rpb246ICdlZGl0JwogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgIGRlbGV0ZTogewogICAgICAgICAgICAgICAgICAgIGNsYXNzOiAnYnRuIGJ0bi1zbSBidG4tZGFuZ2VyJywKICAgICAgICAgICAgICAgICAgICBodG1sOiAnPHNwYW4gY2xhc3M9ImdseXBoaWNvbiBnbHlwaGljb24tcmVtb3ZlIj48L3NwYW4+ICZuYnNwIERFTEVURScsCiAgICAgICAgICAgICAgICAgICAgYWN0aW9uOiAnZGVsZXRlJwogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgIGNvbmZpcm06IHsKICAgICAgICAgICAgICAgICAgICBjbGFzczogJ2J0biBidG4tc20gYnRuLWRlZmF1bHQnLAogICAgICAgICAgICAgICAgICAgIGh0bWw6ICc8c3BhbiBjbGFzcz0iZ2x5cGhpY29uIGdseXBoaWNvbi1vayI+PC9zcGFuPiAmbmJzcCBBcmUgeW91IFN1cmUgPycsCiAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgc2F2ZTogewogICAgICAgICAgICAgICAgICAgIGNsYXNzOiAnYnRuIGJ0bi1zbSBidG4tc3VjY2VzcycsCiAgICAgICAgICAgICAgICAgICAgaHRtbDogJzxzcGFuIGNsYXNzPSJnbHlwaGljb24gZ2x5cGhpY29uLWZsb3BweS1kaXNrIj48L3NwYW4+ICZuYnNwIFNBVkUnCiAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICAgcmVzdG9yZTogewogICAgICAgICAgICAgICAgICAgIGNsYXNzOiAnYnRuIGJ0bi1zbSBidG4td2FybmluZycsCiAgICAgICAgICAgICAgICAgICAgaHRtbDogJzxzcGFuIGNsYXNzPSJnbHlwaGljb24gZ2x5cGhpY29uLXJlcGVhdCI+PC9zcGFuPiAmbmJzcCBSRVNUT1JFJywKICAgICAgICAgICAgICAgICAgICBhY3Rpb246ICdkZWxldGUnCiAgICAgICAgICAgICAgICB9LAogICAgICAgICAgICB9CiAgICAgICAgfSk7CiAgICB9KTsKPC9zY3JpcHQ+Cnt7IGVuZCB9fQp7e3RlbXBsYXRlICJmb290ZXIuaHRtbCIgLn19Cg==",
	"footer.html":         "e3sgZGVmaW5lICJmb290ZXIuaHRtbCIgfX0KICAgICAgICA8L2Rpdj4KICAgIDwvZGl2Pgo8L2Rpdj4gPCEtLSAuL2NvbnRhaW5lciAtLT4KCjxkaXYgY2xhc3M9ImZvb3RlciI+CiAgICAmY29weTsgUm9kbmV5IEFtYXRvIDIwMTYKPC9kaXY+Cgo8c2NyaXB0IHNyYz0iL3B1YmxpYy9qcy9ib290c3RyYXAubWluLmpzIj48L3NjcmlwdD4KCjwvYm9keT4KPC9odG1sPgp7eyBlbmQgfX0=",
	"header.html":         "e3sgZGVmaW5lICJoZWFkZXIuaHRtbCIgfX0KPCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KICAgIDxoZWFkPgogICAgICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgICAgICA8bWV0YSBodHRwLWVxdWl2PSJYLVVBLUNvbXBhdGlibGUiIGNvbnRlbnQ9IklFPWVkZ2UiPgogICAgICAgIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MSI+CgogICAgICAgIDx0aXRsZT57ey5UaXRsZX19IDogTHVuY2htb25leTwvdGl0bGU+CgogICAgICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iL3B1YmxpYy9jc3MvYm9vdHN0cmFwLm1pbi5jc3MiPgogICAgICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iL3B1YmxpYy9wbHVnaW5zL2ZvbnQtYXdlc29tZS9jc3MvZm9udC1hd2Vzb21lLm1pbi5jc3MiPgogICAgICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iL3B1YmxpYy9jc3MvYXBwLmNzcyI+CiAgICAgICAgPHNjcmlwdCBzcmM9Ii9wdWJsaWMvanMvanF1ZXJ5LTIuMS4zLm1pbi5qcyI+PC9zY3JpcHQ+CiAgICA8L2hlYWQ+Cjxib2R5PgoKPGRpdiBjbGFzcz0iY29udGFpbmVyLWZsdWlkIj4KPGRpdiBjbGFzcz0icm93Ij4KICAgIDxkaXYgY2xhc3M9ImNvbC1tZC0xMCBjb2wtbWQtb2Zmc2V0LTIgY29sLXNtLTEyIj4KICAgICAgICA8YSBjbGFzcz0ibmF2YmFyLWJyYW5kIiBocmVmPSIvIj4KICAgICAgICAgIDxzcGFuIGNsYXNzPSJnbHlwaGljb24gZ2x5cGhpY29uLXVzZXIgZ2x5cGhpY29uLXNjaXNzb3JzIj48L3NwYW4+CiAgICAgICAgICBMdW5jaG1vbmV5CiAgICAgICAgPC9hPgogICAgPC9kaXY+CjwvZGl2Pgo8L2Rpdj4KCjxkaXYgY2xhc3M9ImNvbnRhaW5lci1mbHVpZCI+CiAgICA8ZGl2IGNsYXNzPSJyb3ciPgogICAgICAgIDxkaXYgY2xhc3M9ImNvbC1tZC0yIGNvbC1zbS0xMiBzaWRlLW1lbnUtY29udGFpbmVyIj4KICAgICAgICB7e3RlbXBsYXRlICJzaWRlbWVudS5odG1sIiAufX0KICAgICAgICA8L2Rpdj4KICAgICAgICA8ZGl2IGNsYXNzPSJjb2wtbWQtMTAgY29sLXNtLTEyIGNvbnRlbnQiPgp7e2VuZH19",
	"index.html":          "Cnt7dGVtcGxhdGUgImhlYWRlci5odG1sIiAufX0KCjxoMz5XZWxjb21lIHRvIGx1bmNobW9uZXk8L2gzPgoKe3t0ZW1wbGF0ZSAiZm9vdGVyLmh0bWwiIC59fQo=",
	"interest.html":       "Cnt7dGVtcGxhdGUgImhlYWRlci5odG1sIiAufX0KCjxkaXYgY2xhc3M9InJvdyI+CiAgICA8ZGl2IGNsYXNzPSJjb2wtbWQtMTIiPgogICAgICAgIDx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtc3RyaXBlZCB0YWJsZS1ib3JkZXJlZCBlZGl0YWJsZS10YWJsZSI+CiAgICAgICAgPHRoZWFkPgogICAgICAgIDx0cj4KICAgICAgICAgICAgPHRoPkRheTwvdGg+CiAgICAgICAgICAgIDx0aD5CYWxhbmNlPC90aD4KICAgICAgICAgICAgPHRoPkludGVyZXN0PC90aD4KICAgICAgICAgICAgPHRoPkRheSBJbnRlcmVzdDwvdGg+CiAgICAgICAgICAgIDx0aD5SdW5uaW5nIEludGVyZXN0PC90aD4KICAgICAgICA8L3RyPgogICAgICAgIDwvdGhlYWQ+CiAgICAgICAgPHRmb290PgogICAgICAgIDx0cj4KICAgICAgICAgICAgPHRkPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48L3RkPgogICAgICAgICAgICA8dGQ+PC90ZD4KICAgICAgICAgICAgPHRkPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48L3RkPgogICAgICAgIDwvdHI+CiAgICAgICAgPC90Zm9vdD4KICAgICAgICA8dGJvZHk+CiAgICAgICAge3sgcmFuZ2UgLkludGVyZXN0IH19CiAgICAgICAgICAgIDx0cj4KICAgICAgICAgICAgICAgIDx0ZD57eyAuRGF5IH19PC90ZD4KICAgICAgICAgICAgICAgIDx0ZD57eyAuQmFsYW5jZSB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQ+e3sgLkludGVyZXN0UmF0ZSB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQ+e3sgLkRheXNJbnRlcmVzdCB9fTwvdGQ+CiAgICAgICAgICAgICAgICA8dGQ+e3sgLlJ1bm5pbmdJbnRlcmVzdCB9fTwvdGQ+CiAgICAgICAgICAgIDwvdHI+CiAgICAgICAge3sgZW5kIH19CiAgICAgICAgPC90Ym9keT4KICAgICAgICA8L3RhYmxlPgogICAgPC9kaXY+CjwvZGl2PgoKe3t0ZW1wbGF0ZSAiZm9vdGVyLmh0bWwiIC59fQo=",
	"shares_index.html":   "e3t0ZW1wbGF0ZSAiaGVhZGVyLmh0bWwiIC59fQo8dGFibGUgY2xhc3M9InRhYmxlIHRhYmxlLXN0cmlwZWQgdGFibGUtYm9yZGVyZWQiPgo8dGhlYWQ+Cjx0cj4KPHRoPlN0b2NrIENvZGU8L3RoPgo8dGg+UXR5PC90aD4KPHRoPlZhbHVlPC90aD4KPC90cj4KPC90aGVhZD4KPHRib2R5PgoKe3sgcmFuZ2UgLlNoYXJlcyB9fQoKPHRyPgogICAgPHRkPnt7IC5Db2RlIH19PC90ZD4KICAgIDx0ZCBhbGlnbj0icmlnaHQiPnt7IC5RdHkgfX08L3RkPgogICAgPHRkIGFsaWduPSJyaWdodCI+e3sgLlZhbHVlIH19PC90ZD4KPC90cj4KCnt7IGVuZCB9fQo8L3Rib2R5Pgo8L3RhYmxlPgoKe3t0ZW1wbGF0ZSAiZm9vdGVyLmh0bWwiIC59fQo=",
	"sidemenu.html":       "e3sgZGVmaW5lICJzaWRlbWVudS5odG1sIiB9fQoKPHRhYmxlIGNsYXNzPSJ0YWJsZSI+Cjx0aGVhZD4KPC90aGVhZD4KPHRib2R5Pgo8dHI+CiAgICA8dGQ+PGEgaHJlZj0iLyI+PGkgY2xhc3M9ImZhIGZhLWhvbWUgZmEtZnciPjwvaT5Ib21lPC9hPjwvdGQ+CiAgICA8dGQ+PC90ZD4KPC90cj4Ke3sgcmFuZ2UgJGEgOj0gLkFjY291bnRzIH19Cgo8dHI+CiAgICA8dGQ+CiAgICA8YSBocmVmPSIvYWNjb3VudHMve3sgJGEuSUQgfX0iIGNsYXNzPSJhY3RpdmUiPjxpIGNsYXNzPSJmYSBmYS17eyAkYS5JY29uIH19IGZhLWZ3Ij48L2k+IHt7LSAkYS5OYW1lIH19PC9hPgogICAgPC90ZD4KICAgIDx0ZCBjbGFzcz0idGV4dC1yaWdodCI+CiAgICA8YSBocmVmPSIvYWNjb3VudHMve3sgJGEuSUQgfX0iIGNsYXNzPSJhY3RpdmUiPgogICAgPHNwYW4gY2xhc3M9InRvdGFsIj48aSBjbGFzcz0iZmEgZmEtZnciPjwvaT57eyAkYS5HZXRGb3JtYXR0ZWRUb3RhbCB9fTwvc3Bhbj4KICAgIDwvYT4KICAgIDwvdGQ+CjwvdHI+Cnt7IGVuZCB9fQo8dHI+CiAgICA8dGQ+PGEgaHJlZj0iL2ludGVyZXN0P2Zyb209MjAxNi0wMy0yOCZhbXA7dW50aWw9MjAxNi0wNC0yNyI+SW50ZXJlc3Q8L2E+PC90ZD4KPC90cj4KPHRyPgogICAgPHRkPjxhIGhyZWY9Ii9zaGFyZXMiPlNoYXJlczwvYT48L3RkPgo8L3RyPgo8L3Rib2R5Pgo8L3RhYmxlPgp7eyBlbmQgfX0K",
}
