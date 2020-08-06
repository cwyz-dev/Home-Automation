package controller

import
(
	"github.com/cwyz-dev/Home-Automation/config-management-service/domain"
	"net/http"
	"github.com/gorilla/mux"
)

type Controller struct
{
	Config *domain.Config
}

// ReadConfig writes the configuration for the given service to the ResponseWriter
func (c *Controller) ReadConfig(w http.ResponseWriter, r *http.Request)
{
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	vars := mux.Vars(r)
	serviceName,ok := vars["serviceName"]
	if !ok
	{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error")
	}
	config,err := c.Config.Get(serviceName)
	if err != nil
	{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}
	rsp,err := json.Marshal(&config)
	if err != nil
	{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(rsp))
}
