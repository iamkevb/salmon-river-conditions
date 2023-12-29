function createPrecipitationChart() {
  const data = {
    labels: [{{ range .Dates }} "{{.}}",{{end}}],
    datasets: [
      {
        label: "Rain",
        backgroundColor: "rgba(54, 162, 235, 0.4)",
        borderColor: "rgba(54, 162, 235, 1)",
        hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
        borderWidth: 2,
        data: [{{ range .Rain }} "{{.}}",{{end}}],
      },
      {
        label: "Snow",
        backgroundColor: "rgba(255, 99, 132, 0.4)",
        borderColor: "rgba(255, 99, 132, 1)",
        hoverBackgroundColor: "rgba(255, 99, 132, 0.9)",
        borderWidth: 2,
        data: [{{ range .Snow }} "{{.}}",{{end}}],
      },
    ],
  }

  const precipationToolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label 
  }

  const precipationToolTipLabel = (toolTipItem) => {
    pType = toolTipItem.datasetIndex === 0 ? "Rain" : "Snow"
    pUnit = toolTipItem.datasetIndex === 0 ? "mm" : "cm"
    return `${pType}: ${toolTipItem.formattedValue} ${pUnit}`
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
            title: precipationToolTipTitle,
            label: precipationToolTipLabel,
          },
        }
      }
    }
  }

  const precipationCtx = document.getElementById("precipitationChart")
  new Chart(precipationCtx, config)
}

document.addEventListener("DOMContentLoaded", () => {
  createPrecipitationChart()
});