class ChangeTimestampsToDatetimeInCategories < ActiveRecord::Migration[7.2]
  def change
    change_column :categories, :created_at, :datetime
    change_column :categories, :updated_at, :datetime
  end
end
