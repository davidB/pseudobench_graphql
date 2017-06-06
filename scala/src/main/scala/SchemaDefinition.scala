import sangria.execution.deferred.{Fetcher, HasId}
import sangria.schema._

import scala.concurrent.Future

object SchemaDefinition {
  val Query = ObjectType(
    "Query", fields[Object, Unit](
      Field("hello", StringType,
          None,
          resolve = (ctx) ⇒ "world")
    ))

  val MySchema = Schema(Query)
}
