package handlers

import (
	"fmt"
	"net/http"
)

// DefaultHandler handler struct
type DefaultHandler struct{}

//GetDefault - returns defult message
func (dh *DefaultHandler) GetDefault(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<html>
	    <body>
	    <h1>EDO API</h1>
		<h2>Bojanche Stojchevski</h2>
		<p>
			Available endpoints:
			<ul>
				<li><h3>[GET]</h3> '/version' - returns the version and host </li>
				<hr />
				<li><h3>[GET]</h3> '/jobs' - returns all the jobs
				<ul>
				Parameters:
				<li>sortBy=[all available fields for a job]</li>
				</ul>
				</li>
				 
				<li><h3>[POST]</h3> '/jobs'
				<ul>
				Parameters:
				<li>priority=[integer value]</li>
				</ul>
				</li>
		
			</ul>
		</p>
	    </body>
	</html>	
	`)
}
