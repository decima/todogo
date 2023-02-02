package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"net/http"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var taskList = []Task{}

// Pending Tasks metrics Example
var nbPendingTasks = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "nb_pending",
	Help: "Number of hard-disk errors.",
})

func init() {
	prometheus.Register(version.NewCollector("todogo"))

	// Pending Tasks metrics Example
	prometheus.MustRegister(nbPendingTasks)
}

func main() {
	taskList = append(taskList, Task{Description: "Learn Docker", Done: true})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"tasks": taskList,
		})
	})

	r.GET("/toggle", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		taskList[id].Done = !taskList[id].Done

		//Pending task metrics
		if !taskList[id].Done {
			nbPendingTasks.Inc()
		} else {
			nbPendingTasks.Dec()
		}

		c.Redirect(302, "/")
	})

	r.GET("/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))

		//Pending task metrics
		nbPendingTasks.Dec()

		taskList = append(taskList[:id], taskList[id+1:]...)
		c.Redirect(302, "/")

	})

	r.GET("/add", func(c *gin.Context) {
		description := c.Query("description")

		//Pending task metrics
		nbPendingTasks.Inc()

		taskList = append(taskList, Task{Description: description})
		c.Redirect(302, "/")

	})

	r.GET("/metrics", func(c *gin.Context) {
		h := promhttp.HandlerFor(prometheus.Gatherers{
			prometheus.DefaultGatherer,
		}, promhttp.HandlerOpts{})
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run("0.0.0.0:9000")
}
