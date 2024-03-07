use std::{fs::File, io::BufReader, io::BufRead, path::Path};

use actix_web::{web, App, HttpResponse, HttpRequest, HttpServer, Responder, middleware::Logger};
use rand::seq::SliceRandom;
use env_logger;

struct JokeData {
    jokes: Vec<String>,
}

async fn joke_handler(_req: HttpRequest, joke_data: web::Data<JokeData>) -> impl Responder {
    let joke: String = joke_data.jokes.choose(&mut rand::thread_rng()).unwrap().clone();
    println!("joke chosen: {joke}");
    joke
}

fn load_jokes(filepath: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filepath).expect("No file found");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Unable to read line"))
        .collect()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));

    let jokes = load_jokes("/home/rob/Documents/Projects/Micro90s/Micro90s/jokes.txt");

    HttpServer::new(move || App::new()
            .app_data(web::Data::new( JokeData{jokes: jokes.clone()}))
            .wrap(Logger::default())
            .route("/", web::get().to(HttpResponse::Ok))
            .route("/joke", web::get().to(joke_handler))
        )
        .bind(("127.0.0.1", 5050))?
        .run()
        .await
}