package castellancore


const helpMessage = `Thank you for using Castellan! The following commands are currently available:
- **/register [campaign name] <weather system>**: Registers a new campaing. Be sure to give your campaign a name and weather system. Currently, 'Yoon Suin' is the only weather system supported.

- **/weather**: Rolls the next day in the weather chain. Outputs a short synopsis of the results.

- **/weather-report [yyyy-mm-dd] <location>**: Gives a detailed report for the given date. If the date has not been rolled, simply returns an error message. If the weather system has locations configure, you can optionally put in a location name and get on the information for it
	`//- **/schedule-report [hh:mm] [timzone]**: Schedules a report to be delivered at a certain time 