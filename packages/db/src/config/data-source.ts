import { DataSource } from "typeorm";
import * as dotenv from "dotenv";

dotenv.config();
export const AppDataSource = new DataSource({
  type: "postgres",
  url: `postgres://${process.env.DATABASE_USERNAME}:${process.env.DATABASE_PASSWORD}@${process.env.DATABASE_HOST}:${process.env.DOCKER_DATABASE_PORT}/${process.env.DATABASE_NAME}`,
  entities: [],
  migrations: [__dirname + "/../migrations/*.{ts,js}"],
});

AppDataSource.initialize()
  .then(() => {
    console.log("database initialized...");
  })
  .catch((err) => {
    console.log("error in data-source", err);
  });
