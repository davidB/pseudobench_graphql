name := "akka-http-sangria"
version := "0.1.0-SNAPSHOT"

description := "adaptation of sangria-akka-http-example for the helloworld benchmark"

scalaVersion := "2.12.1"
scalacOptions ++= Seq("-deprecation", "-feature")

libraryDependencies ++= Seq(
  "org.sangria-graphql" %% "sangria" % "1.1.0",
  "org.sangria-graphql" %% "sangria-spray-json" % "1.0.0",
  "com.typesafe.akka" %% "akka-http" % "10.0.1",
  "com.typesafe.akka" %% "akka-http-spray-json" % "10.0.1",
)

Revolver.settings
enablePlugins(JavaAppPackaging)
