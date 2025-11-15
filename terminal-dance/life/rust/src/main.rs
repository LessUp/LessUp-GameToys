use life::{config::Config, run};

fn main() {
    let cfg = Config::parse().unwrap_or_else(|err| err.exit());
    if let Err(err) = run(cfg) {
        eprintln!("{err}");
        std::process::exit(1);
    }
}
