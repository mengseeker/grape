class CreateCluster < ActiveRecord::Migration[6.1]
  def change
    create_table :clusters do |t|
      t.string :name, null: false
      t.string :code, null: false
      t.string :deploy_type
      t.string :note

      t.index :code, unique: true
    end
  end
end
