class ChangeDatetimeToTimestampsInCategories < ActiveRecord::Migration[7.2]
  def change
    change_column :categories, :created_at, :timestamp
    change_column :categories, :updated_at, :timestamp
  end
end
