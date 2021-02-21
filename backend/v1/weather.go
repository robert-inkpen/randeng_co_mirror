package v1

// WeatherService interface
type WeatherService interface {
	CurrentAndForecast()
}

type weatherService struct {}

// NewWeatherService initializes and allocates new weather service interface implementation
func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (s *weatherService) CurrentAndForecast() {

}