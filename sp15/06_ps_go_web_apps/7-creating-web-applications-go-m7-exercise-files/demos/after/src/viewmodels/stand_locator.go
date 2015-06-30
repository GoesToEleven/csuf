package viewmodels

import ()

type standLocator struct {
	Title  string
	Active string
}

func GetStandLocator() standLocator {
	result := standLocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "stand_locator",
	}

	return result
}

type standLocation struct {
	Lat   float32
	Lng   float32
	Title string
}

func GetStandLocations() []standLocation {
	result := []standLocation{
		standLocation{
			Lat:   37.4217,
			Lng:   -122.075,
			Title: "Matthew's stand",
		},
		standLocation{
			Lat:   37.4206,
			Lng:   -122.08,
			Title: "Alice's stand",
		},
		standLocation{
			Lat:   37.4205,
			Lng:   -122.083,
			Title: "Kara's stand",
		},
		standLocation{
			Lat:   37.414,
			Lng:   -122.09,
			Title: "Fred's stand",
		},
		standLocation{
			Lat:   37.412,
			Lng:   -122.09,
			Title: "Jake's stand",
		},
		standLocation{
			Lat:   37.41,
			Lng:   -122.093,
			Title: "Wallace's stand",
		},
		standLocation{
			Lat:   37.416,
			Lng:   -122.095,
			Title: "Gromit's stand",
		},
		standLocation{
			Lat:   37.41,
			Lng:   -122.1,
			Title: "Kirk's stand",
		},
		standLocation{
			Lat:   37.41,
			Lng:   -122.102,
			Title: "Lorelei's stand",
		},
		standLocation{
			Lat:   37.412,
			Lng:   -122.099,
			Title: "Rebecca's stand",
		},
		standLocation{
			Lat:   37.407,
			Lng:   -122.1025,
			Title: "Chris's stand",
		},
		standLocation{
			Lat:   37.423,
			Lng:   -122.1025,
			Title: "Carson's stand",
		},
	}
	
	return result
}
