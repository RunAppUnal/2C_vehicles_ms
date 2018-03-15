class CreateVehicles < ActiveRecord::Migration[5.1]
  def change
    create_table :vehicles do |t|
      t.string :plate
      t.integer :user_id
      t.string :kind
      t.integer :model
      t.string :color
      t.integer :capacity
      t.string :image
      t.string :brand

      t.timestamps
    end
  end
end
