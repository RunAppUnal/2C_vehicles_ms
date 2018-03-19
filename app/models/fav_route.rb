class FavRoute < ApplicationRecord
	before_save :update_count

	validates :user_id, presence: true
    validates :count, numericality: { :less_than_or_equal_to => 5, message: 'Solo se pueden agregar 5 rutas favoritas.' }

private

  def update_count
  	self.count = 0
    self.count += 1 if self.polyline1.present?
    self.count += 1 if self.polyline2.present?
    self.count += 1 if self.polyline3.present?
    self.count += 1 if self.polyline4.present?
    self.count += 1 if self.polyline5.present?
  end
end