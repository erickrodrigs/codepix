import { Command, Console } from 'nestjs-console';
import { getConnection } from 'typeorm';
import fixtures from './fixtures';
import * as chalk from 'chalk';

@Console()
export class FixturesCommand {
  @Command({
    command: 'fixtures',
    description: 'Seed data in database',
  })
  async command() {
    await this.runMigrations();

    for (const fixture of fixtures) {
      await this.createData(fixture.model, fixture.fields);
    }

    console.log(chalk.green('Data generated'));
  }

  async runMigrations() {
    const connection = getConnection('default');

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    for (const migration of connection.migrations.reverse()) {
      await connection.undoLastMigration();
    }
  }

  async createData(model: any, data: any) {
    const repository = this.getRepository(model);
    const obj = repository.create(data);
    await repository.save(obj);
  }

  getRepository(model: any) {
    const connection = getConnection('default');
    return connection.getRepository(model);
  }
}
