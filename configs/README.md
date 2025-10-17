# Configuration Directory

This directory contains configuration files for the application.

## Files

- `config.yaml` - Default configuration file (tracked in git)
- `config.example.yaml` - Example configuration template
- `config.local.yaml` - Local overrides (not tracked in git)
- `config.*.yaml` - Environment-specific configs (not tracked in git)

## Usage

### Default Configuration

The application will automatically load `config.yaml` from this directory.

```bash
./build/server
```

### Custom Configuration File

You can specify a custom configuration file:

```go
cfg, err := config.Load("./configs/config.production.yaml")
```

### Environment Variables

You can override any configuration value using environment variables with the `APP_` prefix:

```bash
# Override server settings
export APP_SERVER_PORT=9090
export APP_SERVER_MODE=release

# Override database settings
export APP_DATABASE_HOST=db.example.com
export APP_DATABASE_USER=myuser
export APP_DATABASE_PASSWORD=secret

./build/server
```

Environment variable format: `APP_<SECTION>_<KEY>` (e.g., `APP_SERVER_PORT`, `APP_DATABASE_NAME`)

### Local Development

For local development, create a `config.local.yaml` file (not tracked by git):

```bash
cp config.example.yaml config.local.yaml
# Edit config.local.yaml with your local settings
```

## Configuration Structure

```yaml
server:
  port: "8080"           # Server port
  mode: "debug"          # Gin mode: debug, release, or test

database:
  host: "localhost"      # Database host
  port: "3306"          # Database port
  user: "root"          # Database user
  password: "password"   # Database password
  name: "example_db"    # Database name
```

## Priority

Configuration is loaded in the following order (later overrides earlier):

1. Default values (in code)
2. Configuration file (`config.yaml`)
3. Environment variables (`APP_*`)
