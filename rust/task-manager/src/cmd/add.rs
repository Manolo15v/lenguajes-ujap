use seahorse::{error::FlagError, Command, Context, Flag, FlagType};

fn add_command() {
    Command::new("add")
        .description("add new task")
        .usage("task add [name] ")
        .action(add_action)
}

fn add_action() {
    //TODO make db add logic
}
