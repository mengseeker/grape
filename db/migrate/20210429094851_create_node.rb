class CreateNode < ActiveRecord::Migration[6.1]
  def change
    create_table :nodes do |t|
      t.string :name, null: false
      t.string :code, null: false
      t.integer :namespace_id, null: false
      t.integer :service_id, null: false
      t.integer :cluster_id, null: false
      t.integer :group_id, null: false

      t.string :ip, null: false
      t.integer :state, null: false, default: 0

      t.index :code, unique: true
      t.index :namespace_id
      t.index :cluster_id
      t.index :service_id
      t.index :group_id
    end
  end
end
