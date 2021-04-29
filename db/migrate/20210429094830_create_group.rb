class CreateGroup < ActiveRecord::Migration[6.1]
  def change
    create_table :groups do |t|
      t.string :name, null: false
      t.string :code, null: false
      t.string :version, null: false, default: ""
      t.string :lang, null: false, default: ""
      t.integer :namespace_id, null: false
      t.integer :service_id, null: false
      t.integer :cluster_id, null: false
      t.integer :replicas, default: 0, null: false
      t.integer :deploy_type, default: 0, null: false
      t.string :note, null: false, default: ""

      t.index :code, unique: true
      t.index :namespace_id
      t.index :cluster_id
      t.index :service_id
    end

  end
end
