class CreateGroup < ActiveRecord::Migration[6.1]
  def change
    create_table :groups do |t|
      t.string :name, null: false
      t.string :code, null: false
      t.string :version
      t.string :lang
      t.integer :namespace_id
      t.integer :service_id, null: false
      t.integer :cluster_id, null: false
      t.integer :replicas, default: 0
      t.integer :deploy_type, default: 0
      t.string :note

      t.index :code, unique: true
      t.index :namespace_id
      t.index :cluster_id
      t.index :service_id
    end

  end
end
