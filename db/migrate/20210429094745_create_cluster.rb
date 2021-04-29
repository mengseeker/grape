class CreateCluster < ActiveRecord::Migration[6.1]
  def change
    create_table :clusters do |t|
      t.string :name, null: false
      t.string :code, null: false
      t.string :deploy_type, null: false, default: "k8s"
      t.string :note, null: false, default: ""

      t.index :code, unique: true
    end
  end
end
