class CreateService < ActiveRecord::Migration[6.1]
  def change
    create_table :services do |t|
      t.string :name, null: false
      t.string :code, null: false

      t.integer :port
      t.integer :protocol
      t.integer :external, default: 0

      t.string :note

      t.index :code, unique: true
    end
  end
end
