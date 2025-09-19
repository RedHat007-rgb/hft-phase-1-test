import { CreateDateColumn } from "typeorm";

export class BaseEntity {
  @CreateDateColumn({ type: "timestamptz" })
  createdAt: Date;

  @CreateDateColumn({ type: "timestamptz" })
  updateAt: Date;
}
