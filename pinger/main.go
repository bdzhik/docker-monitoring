package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type PingData struct {
	IP         string    `json:"ip"`
	Time       int       `json:"time"`
	LastActive time.Time `json:"last_active"`
}

func getDockerIPs() []string {
	cmd := exec.Command("sh", "-c", "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -q)")
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error fetching Docker IPs:", err)
		return nil
	}

	ips := strings.Fields(string(out))
	return ips
}

func ping(ip string) int {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	out, err := cmd.Output()
	if err != nil {
		return -1
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		parts := strings.Split(lines[1], "time=")
		if len(parts) > 1 {
			timeMs := strings.Split(parts[1], " ")
			if len(timeMs) > 0 {
				return int(timeMs[0][0])
			}
		}
	}

	return -1
}

func reportPing(ip string) {
	pingTime := ping(ip)
	data := PingData{
		IP:         ip,
		Time:       pingTime,
		LastActive: time.Now(),
	}

	jsonData, _ := json.Marshal(data)
	http.Post("http://backend:8080/ping", "application/json", bytes.NewBuffer(jsonData))
}

func main() {
	for {
		ips := getDockerIPs()
		for _, ip := range ips {
			reportPing(ip)
		}
		time.Sleep(10 * time.Second)
	}
}
