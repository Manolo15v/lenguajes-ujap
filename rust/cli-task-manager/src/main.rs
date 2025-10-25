use seahorse::{error::FlagError, App, Command, Context, Flag, FlagType};
use std::env;

//this is an example from the seahorse docs

fn main() {
 

    app.run(args);
}

fn default_action(c: &Context) {
    if c.bool_flag("bye") {
        println!("Bye, {:?}", c.args);
    } else {
        println!("Hello, {:?}", c.args);
    }

    if let Ok(age) = c.int_flag("age") {
        println!("{:?} is {} years old", c.args, age);
    }
}

fn calc_action(c: &Context) {
    match c.string_flag("operator") {
        Ok(op) => {
            let sum: i32 = match &*op {
                "add" => c.args.iter().map(|n| n.parse::<i32>().unwrap()).sum(),
                "sub" => c.args.iter().map(|n| n.parse::<i32>().unwrap() * -1).sum(),
                _ => panic!("undefined operator..."),
            };

            println!("{}", sum);
        }
        Err(e) => match e {
            FlagError::Undefined => panic!("undefined operator..."),
            FlagError::ArgumentError => panic!("argument error..."),
            FlagError::NotFound => panic!("not found flag..."),
            FlagError::ValueTypeError => panic!("value type mismatch..."),
            FlagError::TypeError => panic!("flag type mismatch..."),
        },
    }
}

fn calc_command() -> Command {
    Command::new("calc")
        .description("calc command")
        .alias("cl, c")
        .usage("cli calc(cl, c) [nums...]")
        .action(calc_action)
        .flag(
            Flag::new("operator", FlagType::String)
                .description("Operator flag(ex. cli calc --operator add 1 2 3)")
                .alias("op"),
        )
}
