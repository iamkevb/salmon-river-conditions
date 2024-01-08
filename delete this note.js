// Fix line graphs to be multiple colours with 2 data sets.


const labels = Utils.months({count: 7});
const data = {
  labels: labels,
  datasets: [
    {
      label: 'Historic',
      data: [0,1,5,3,null,null,null],
      borderColor: Utils.CHART_COLORS.yellow,
      backgroundColor: Utils.transparentize(Utils.CHART_COLORS.red, 0.5),
      tension: 0.2
    },
    {
      label: 'Forecast',
      data: [null,null,null,3,2,1,6],
      borderColor: Utils.CHART_COLORS.green,
      backgroundColor: Utils.transparentize(Utils.CHART_COLORS.blue, 0.5),
      tension: 0.2
    }
  ]
};