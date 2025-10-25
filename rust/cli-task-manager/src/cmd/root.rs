pub let app = App::new(env!("CARGO_PKG_NAME"))
.description(env!("CARGO_PKG_DESCRIPTION"))
.author(env!("CARGO_PKG_AUTHORS"))
.version(env!("CARGO_PKG_VERSION"))
.usage("cli [name]")
.action(default_action)
.flag(
    Flag::new("bye", FlagType::Bool)
        .description("Bye flag")
        .alias("b"),
)
.flag(
    Flag::new("age", FlagType::Int)
        .description("Age flag")
        .alias("a"),
)
.command(calc_command());
let args: Vec<String> = env::args().collect();