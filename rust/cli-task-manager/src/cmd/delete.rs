use seahorse::{error::FlagError, Command, Context, Flag, FlagType};

fn delete_command() {
    Command::new("delete")
        .description("delete new task")
        .usage("task delete [name] ")
        .action(delete_action)
}

fn delete_action() {
    //TODO make db delete logic
}
