class CreateNamespace < ActiveRecord::Migration[6.1]
  def change
    create_table :namespaces do |t|
      t.string :code, null: false
      t.string :name, null: false
      t.string :note

      t.index :code, unique: true
    end
    
  end
end
