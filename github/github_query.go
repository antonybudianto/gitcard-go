package github

import "fmt"

const topSummaryFirst = 10

// UserQuery = query used when fetching profile
var UserQuery = `
query getUserRepo($username: String!, $after: String) {
	user(login:$username){
	  avatarUrl
	  repositories(after:$after, first:100, ownerAffiliations:OWNER, isFork:false, privacy:PUBLIC){
		totalCount
		pageInfo{
		  endCursor
		  hasNextPage
		}
		edges{
		  node{
			name
			forkCount
			primaryLanguage {
			  name
			}
			stargazers {
			  totalCount
			}
		  }
		}
	  }
	}
}
`

func generateSummaryQuery(name, location, language, followers string, first int) string {
	return fmt.Sprintf(`
	%s: search(query: "location:%s language:%s followers:%s", type: USER, first: %d) {
		edges {
			node {
			... on User {
				name
				avatarUrl
				login
				bio
				company
				location
				following {
				totalCount
				}
				followers {
				totalCount
				}
			}
			}
		}
	}
	`, name, location, language, followers, first)
}

// SummaryQuery = query used when fetch all summary
var SummaryQuery = fmt.Sprintf(`
query topSummary {
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
	%s
  }
`,
	generateSummaryQuery("topPHPDev", "Indonesia", "PHP", ">=200", topSummaryFirst),
	generateSummaryQuery("topJsDev", "Indonesia", "JavaScript", ">=200", topSummaryFirst),
	generateSummaryQuery("topJavaDev", "Indonesia", "Java", ">=200", topSummaryFirst),
	generateSummaryQuery("topPythonDev", "Indonesia", "Python", ">=150", topSummaryFirst),
	generateSummaryQuery("topHTMLDev", "Indonesia", "HTML", ">=150", topSummaryFirst),
	generateSummaryQuery("topGoDev", "Indonesia", "Go", ">=100", topSummaryFirst),
	generateSummaryQuery("topRubyDev", "Indonesia", "Ruby", ">=100", topSummaryFirst),
	generateSummaryQuery("topShellDev", "Indonesia", "Shell", ">=100", topSummaryFirst),
	generateSummaryQuery("topSwiftDev", "Indonesia", "Swift", ">=50", topSummaryFirst),

	generateSummaryQuery("topJakartaDev", "Jakarta", "*", ">=300", topSummaryFirst),
	generateSummaryQuery("topBandungDev", "Bandung", "*", ">=200", topSummaryFirst),
	generateSummaryQuery("topYogyakartaDev", "Yogyakarta", "*", ">=100", topSummaryFirst),
	generateSummaryQuery("topMalangDev", "Malang", "*", ">=100", topSummaryFirst),
	generateSummaryQuery("topBaliDev", "Bali", "*", ">=100", topSummaryFirst),
	generateSummaryQuery("topSurabayaDev", "Surabaya", "*", ">=100", topSummaryFirst),
	generateSummaryQuery("topSemarangDev", "Semarang", "*", ">=100", topSummaryFirst),
)

// TopIndonesiaQuery = query for top overall Indonesia, used for star rating
var TopIndonesiaQuery = fmt.Sprintf(`
query topIndonesia {
	%s
  }
`,
	generateSummaryQuery("topIndonesiaDev", "Indonesia", "*", ">=100", 100),
)
