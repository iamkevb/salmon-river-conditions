const precipitationData = {
    datasets: [
      {
        backgroundColor: "rgba(54, 162, 235, 0.4)",
        borderColor: "rgba(54, 162, 235, 1)",
        hoverBackgroundColor: "rgba(54, 162, 235, 0.9)",
        borderWidth: 2,
        borderSkipped: false,
        data: [65, 59, 80, 81, 56, 55, 40],
      },
    ],
  }
  
  const precipCtx = document.getElementById("precipitationChart")
  
//   const toolTipTitle = (toolTipItems) => {
//     return toolTipItems[0].label 
//   }
  
//   const toolTipLabel = (toolTipItem) => {
//     return [`Rain: ${toolTipItem.raw.y[1]}`, `Snow: ${toolTipItem.raw.y[0]}`]
//   }
  
  const precipitationChart = new Chart(precipCtx, {
    type: "bar",
    data: precipitationData,
    options: {
      plugins: {
        legend: false,
        tooltip: {
          displayColors: false,
        //   callbacks: {
        //     title: toolTipTitle,
        //     label: toolTipLabel,
        //   },
        },
      },
    },
  });
  