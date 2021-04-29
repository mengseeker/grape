require 'rack'
require 'active_record'
require 'active_support/all'

include ActiveRecord::Tasks

class Seeder
  def initialize(seed_file)
    @seed_file = seed_file
  end

  def load_seed
    raise "Seed file '#{@seed_file}' does not exist" unless File.file?(@seed_file)
    load @seed_file
  end
end

root = Dir.pwd
DatabaseTasks.env = 'production'
DatabaseTasks.database_configuration =  YAML.load_file("config/database.yaml")
DatabaseTasks.db_dir = File.join root, 'db'
DatabaseTasks.fixtures_path = File.join root, 'test/fixtures'
DatabaseTasks.migrations_paths = [File.join(root, 'db/migrate')]
DatabaseTasks.seed_loader = Seeder.new File.join root, 'db/seeds.rb'
DatabaseTasks.root = root

task :environment do
  ActiveRecord::Base.configurations = DatabaseTasks.database_configuration
  ActiveRecord::Base.establish_connection DatabaseTasks.env.to_sym
end


task :g, [:name] do
  require 'rails/all'
  require 'rails/generators'
  require 'rails/generators/active_record/migration/migration_generator'
  g = ActiveRecord::Generators::MigrationGenerator.new [ENV['name']]
  g.create_migration_file
end

load 'active_record/railties/databases.rake'