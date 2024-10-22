class Post < ApplicationRecord
  before_save :generate_slug
  belongs_to :category, optional: true
  validates :author, presence: true

  def generate_slug
    self.slug = title.parameterize if slug.blank?
  end
end
