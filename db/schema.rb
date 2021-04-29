# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2021_04_29_094851) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "clusters", force: :cascade do |t|
    t.string "name", null: false
    t.string "code", null: false
    t.string "deploy_type"
    t.string "note"
    t.index ["code"], name: "index_clusters_on_code", unique: true
  end

  create_table "groups", force: :cascade do |t|
    t.string "name", null: false
    t.string "code", null: false
    t.string "version"
    t.string "lang"
    t.integer "namespace_id"
    t.integer "service_id", null: false
    t.integer "cluster_id", null: false
    t.integer "replicas", default: 0
    t.integer "deploy_type", default: 0
    t.string "note"
    t.index ["cluster_id"], name: "index_groups_on_cluster_id"
    t.index ["code"], name: "index_groups_on_code", unique: true
    t.index ["namespace_id"], name: "index_groups_on_namespace_id"
    t.index ["service_id"], name: "index_groups_on_service_id"
  end

  create_table "migrations", id: :text, force: :cascade do |t|
    t.datetime "applied_at"
  end

  create_table "namespaces", force: :cascade do |t|
    t.string "code", null: false
    t.string "name", null: false
    t.string "note"
    t.index ["code"], name: "index_namespaces_on_code", unique: true
  end

  create_table "nodes", force: :cascade do |t|
    t.string "name", null: false
    t.string "code", null: false
    t.integer "namespace_id", null: false
    t.integer "service_id", null: false
    t.integer "cluster_id", null: false
    t.integer "group_id", null: false
    t.string "ip", null: false
    t.string "state"
    t.index ["cluster_id"], name: "index_nodes_on_cluster_id"
    t.index ["code"], name: "index_nodes_on_code", unique: true
    t.index ["group_id"], name: "index_nodes_on_group_id"
    t.index ["namespace_id"], name: "index_nodes_on_namespace_id"
    t.index ["service_id"], name: "index_nodes_on_service_id"
  end

  create_table "services", force: :cascade do |t|
    t.string "name", null: false
    t.string "code", null: false
    t.integer "port"
    t.integer "protocol"
    t.integer "external", default: 0
    t.string "note"
    t.index ["code"], name: "index_services_on_code", unique: true
  end

end
