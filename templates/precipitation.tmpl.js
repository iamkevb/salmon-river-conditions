const precipationData = {
  labels: ["Day 1", "Day 2", "Day 3", "Day 4", "Day 5"],
  datasets: [
    {
      label: "Rain",
      backgroundColor: "rgba(54, 162, 235, 0.5)",
      borderColor: "rgba(54, 162, 235, 1)",
      borderWidth: 1,
      data: [20, 22, 19, 21, 18],
    },
    {
      label: "Snow",
      backgroundColor: "rgba(255, 99, 132, 0.5)",
      borderColor: "rgba(255, 99, 132, 1)",
      borderWidth: 1,
      data: [30, 32, 29, 31, 28],
    },
  ],
}

const precipationCtx = document.getElementById("precipitationChart")

const precipationToolTipTitle = (toolTipItems) => {
  return toolTipItems[0].label 
}

const precipationToolTipLabel = (toolTipItem) => {
  pType = toolTipItem.datasetIndex === 0 ? "Rain" : "Snow"
  return `${pType}: ${toolTipItem.formattedValue} mm`
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
