#![allow(unused_imports)]
extern crate iron;
extern crate mount;
extern crate logger;
#[macro_use] extern crate juniper;

use std::env;

use mount::Mount;
use logger::Logger;
use iron::prelude::*;
use juniper::EmptyMutation;
use juniper::iron_handlers::{GraphQLHandler, GraphiQLHandler};
use juniper::FieldResult;

struct QueryRoot;
struct Database {}

// The context object is passed down to all referenced types - all your exposed
// types need to have the same context type.
graphql_object!(QueryRoot: Database as "Query" |&self| {

    field hello() -> String {
        //executor.context().users.get(&id)
        "world".to_string()
    }
});

fn context_factory(_: &mut Request) -> Database {
    Database{}
}

fn main() {
    let mut mount = Mount::new();

    let graphql_endpoint = GraphQLHandler::new(
        context_factory,
        QueryRoot,
        EmptyMutation::<Database>::new(),
    );
    //let graphiql_endpoint = GraphiQLHandler::new("/graphiql");

    //mount.mount("/", graphiql_endpoint);
    mount.mount("/graphql", graphql_endpoint);

    //let (logger_before, logger_after) = Logger::new(None);

    let chain = Chain::new(mount);
    //let mut chain = Chain::new(mount);
    //chain.link_before(logger_before);
    //chain.link_after(logger_after);

    let host = env::var("LISTEN").unwrap_or("0.0.0.0:3003".to_owned());
    println!("GraphQL server started on {}", host);
    Iron::new(chain).http(host.as_str()).unwrap();
}
