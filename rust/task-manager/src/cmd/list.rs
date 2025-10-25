use seahorse::{error::FlagError, Command, Context, Flag, FlagType};

fn list_command() {
    Command::new("list")
        .description("list new task")
        .usage("task list [name] ")
        .action(list_action)
}

fn list_action() {
    //TODO make db list logic
}
