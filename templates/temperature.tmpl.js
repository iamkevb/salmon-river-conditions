const temperatureData = {
  datasets: [
    {
      backgroundColor: "rgba(54, 162, 235, 0.4)",
      borderColor: "rgba(54, 162, 235, 1)",
      hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
      borderWidth: 2,
      borderSkipped: false,
      data: {{.}},
    },
  ],
}

const ctx = document.getElementById("temperatureChart");

const toolTipTitle = (toolTipItems) => {
  return toolTipItems[0].label 
}

const toolTipLabel = (toolTipItem) => {
  return [`High: ${toolTipItem.raw.y[1]}`, `Low: ${toolTipItem.raw.y[0]}`]
}

const temperatureChart = new Chart(ctx, {
  type: "bar",
  data: temperatureData,
  options: {
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
});
