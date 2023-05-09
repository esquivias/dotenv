# dotenv

`dotenv` is a small command-line utility that helps you run commands with environment variables loaded from a `.env` file.

## Usage

```
dotenv [[-f <name>][-s <separator>]] <command> [<args>...]
```

### Options

- `-f <name>`
  - Specify the name of the environment file to load (defaults to `.env`).
- `-s <separator>`
  - Specify the key-value separator to use (defaults to `=`).

The remaining arguments are treated as the command to run and its arguments.

## Example

```
./dotenv -f local.env ./server -port 8080
```

This will load the environment variables from the `local.env` file, and then run the `server` command with the `port` argument.

## Building from Source

You can build `dotenv` from source by cloning this repository and running `go build`.

```
git clone https://github.com/esquivias/dotenv.git
cd dotenv
go build -o dotenv
```

This will produce an executable named `dotenv` in the current directory.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [license](license) file for details.
