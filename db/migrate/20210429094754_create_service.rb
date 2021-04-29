class CreateService < ActiveRecord::Migration[6.1]
  def change
    create_table :services do |t|
      t.string :name, null: false
      t.string :code, null: false

      t.integer :port, null: false
      t.integer :protocol, null: false
      t.integer :external, null: false, default: 0

      t.string :note, null: false, default: ""

      t.index :code, unique: true
    end
  end
end
