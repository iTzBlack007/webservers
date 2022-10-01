#[macro_use]
extern crate rocket;

#[get("/")]
fn say_hello() -> &'static str {
    "Hello world"
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![say_hello])
}
