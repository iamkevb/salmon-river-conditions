const precipationData = {
  labels: [{{ range .Dates }} "{{.}}",{{end}}],
  datasets: [
    {
      label: "Rain",
      backgroundColor: "rgba(54, 162, 235, 0.5)",
      borderColor: "rgba(54, 162, 235, 1)",
      borderWidth: 1,
      data: [{{ range .Rain }} "{{.}}",{{end}}],
    },
    {
      label: "Snow",
      backgroundColor: "rgba(255, 99, 132, 0.5)",
      borderColor: "rgba(255, 99, 132, 1)",
      borderWidth: 1,
      data: [{{ range .Snow }} "{{.}}",{{end}}],
    },
  ],
}

const precipationCtx = document.getElementById("precipitationChart")

const precipationToolTipTitle = (toolTipItems) => {
  return toolTipItems[0].label 
}

const precipationToolTipLabel = (toolTipItem) => {
  pType = toolTipItem.datasetIndex === 0 ? "Rain" : "Snow"
  pUnit = toolTipItem.datasetIndex === 0 ? "mm" : "cm"
  return `${pType}: ${toolTipItem.formattedValue} ${pUnit}`
}

const precipitationChart = new Chart(precipationCtx, {
  type: "bar",
  data: precipationData,
  options: {
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
});
