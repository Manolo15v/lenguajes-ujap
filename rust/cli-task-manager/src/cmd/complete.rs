use seahorse::{error::FlagError, Command, Context, Flag, FlagType};

fn complete_command() {
    Command::new("complete")
        .description("complete new task")
        .usage("task complete [name] ")
        .action(complete_action)
}

fn complete_action() {
    //TODO make db complete logic
}
