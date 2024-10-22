class AddAuthorToPosts < ActiveRecord::Migration[7.2]
  def change
    add_column :posts, :author, :string
  end
end
