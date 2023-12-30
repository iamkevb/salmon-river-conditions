# To Do

- [x] Intoduce .env for development and production
- [x] .dev should load sample.json from file file
- [x] .dev should set cache headers to no cache
- [x] .prod shouls set cache headers to 1 hour
- [x] Add precipation graph (total rain, snow by day)
- [x] Sync fetch to prevent more than one at a time
- [x] Add pressure graph (hourly)
- [x] Can temperature graph work similar to precipitation?
- [x] Move temperature and precipitation to functions
- [x] Add water level graph
- [x] Make responsive for phones
- [x] Deploy
- [x] Github action to deploy on merge to main
- [ ] Use USGS as param, load weather with lat/lon from response
- [ ] Combine `SiteData` and `WeatherData`
- [ ] Cache data together
- [ ] Invent some extraData.json file that let me put links or render extra template for a usgs number
- [ ] Different colors for historical and forecasted data
- [ ] Better titles 
    - [ ] Subtitle can be lighter grey
    - [ ] Total rainfall in last x days: 3mm 
    - [ ] Total snowfall in last x days: 3cm
    - [ ] Max and min temperature
    - [ ] Forecast air pressure prediction, or current trend

    Can I get Salmon River Reservoir data?
    `https://api.safewaters.com/api/schedule/3524dbf0-00c8-11ec-9351-dd66b05aaa5c`