class CreateFavRoutes < ActiveRecord::Migration[5.1]
  def change
    create_table :fav_routes do |t|
      t.integer :user_id
      t.string :polyline1
      t.string :polyline2
      t.string :polyline3
      t.string :polyline4
      t.string :polyline5
      t.integer :count

      t.timestamps
    end
  end
end
