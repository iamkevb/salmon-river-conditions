function createTemperatureChart() {
  const data = {
    labels: [
      {{ range .Dates }} 
        "{{.}}", 
      {{end}}
    ],
    datasets: [
      {
        backgroundColor: "rgba(54, 162, 235, 0.4)",
        borderColor: "rgba(54, 162, 235, 1)",
        hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
        borderWidth: 2,
        borderSkipped: false,
        data: [{{ range .Temps }} 
          [{{ range . }} {{.}}, {{end}}], 
        {{end}}],
      },
    ],
  }

  
  const toolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label 
  }

  const toolTipLabel = (toolTipItem) => {
    return [`High: ${toolTipItem.raw[1]}`, `Low: ${toolTipItem.raw[0]}`]
  }

  const config = {
    type: "bar",
    data: data,
    options: {
      maintainAspectRatio: false,
      plugins: {
        legend: false,
        tooltip: {
          displayColors: false,
          callbacks: {
            title: toolTipTitle,
            label: toolTipLabel,
          },
        },
      },
    },
  }

  const ctx = document.getElementById("temperatureChart")
  const temperatureChart = new Chart(ctx, config)
}

document.addEventListener("DOMContentLoaded", () => {
  createTemperatureChart()
});
