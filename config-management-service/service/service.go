package service

import
(
	"time"
)

type ConfigService struct
{
	Config *domain.Config
	Location string
}

// Watch reloads the config every d duration
func (s *ConfigService) Watch(d time.Duration)
{
	for
	{
		err := s.Reload()
		if err != nil
		{
			log.Print(err)
		}
		time.Sleep(d)
	}
}

// Reload reads the configuration and applies changes 
func (s *configService) Reload() error
{
	data, err := ioutil.ReadFile(s.Location)
	if err != nil
	{
		return err
	}
	err = s.Config.SetFromBytes(data)
	if err != nil
	{
		return err
	}
	return nil
}
