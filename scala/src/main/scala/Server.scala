import akka.actor.ActorSystem
import akka.http.scaladsl.Http
import akka.http.scaladsl.server.Directives._
import akka.http.scaladsl.model.StatusCodes._
import akka.http.scaladsl.server._
import akka.stream.ActorMaterializer
import akka.http.scaladsl.marshallers.sprayjson.SprayJsonSupport._
import sangria.execution.deferred.DeferredResolver
import sangria.parser.QueryParser
import sangria.execution.{ErrorWithResolver, Executor, QueryAnalysisError}
import sangria.marshalling.sprayJson._
import spray.json._

import scala.util.{Failure, Success}

object Server extends App {
  implicit val system = ActorSystem("sangria-server")
  implicit val materializer = ActorMaterializer()

  import system.dispatcher

  val route: Route =
    (get & path("graphql")) {
      parameters('query.as[String], 'operation.as[String].?) { (query, operation) ⇒
        val vars = JsObject.empty

        QueryParser.parse(query) match {

          // query parsed successfully, time to execute it!
          case Success(queryAst) ⇒
            complete(Executor.execute(SchemaDefinition.MySchema, queryAst, MyData,
                variables = vars,
                operationName = operation
                //deferredResolver = DeferredResolver.fetchers(SchemaDefinition.characters)
                )
              .map(OK → _)
              .recover {
                case error: QueryAnalysisError ⇒ BadRequest → error.resolveError
                case error: ErrorWithResolver ⇒ InternalServerError → error.resolveError
              })

          // can't parse GraphQL query, return error
          case Failure(error) ⇒
            complete(BadRequest, JsObject("error" → JsString(error.getMessage)))
        }
      }
    }


  Http().bindAndHandle(route, "0.0.0.0", sys.props.get("http.port").fold(3003)(_.toInt))
}
