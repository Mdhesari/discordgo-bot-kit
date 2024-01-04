package config

type Discord struct {
	Name    string `koanf:"name"`
	GuildID string `koanf:"guild_id"`
	Token   string `koanf:"token"`
}

type Config struct {
	Discord Discord `koanf:"discord"`
}

func New(discord Discord) Config {
	return Config{
		Discord: discord,
	}
}
