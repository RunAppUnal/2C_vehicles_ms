class Vehicle < ApplicationRecord
	#----------------------- Validations for vehicles ----------------------
    validates :plate, :kind, :model, :brand, :color, :capacity, presence: true
    validates :model, :capacity, :numericality => {:only_integer => true, message: "debe ser numérico."}
    validates_format_of :plate, with: /[A-Z][A-Z][A-Z]-[0-9][0-9][0-9]/, message: 'debe tener un formato como AAA-111 (sin espacios y en mayúsculas).', 
                            :if => "kind == 'Carro'"
    validates_format_of :plate, with: /[A-Z][A-Z][A-Z]-[0-9][0-9][A-Z]/, message: 'debe tener un formato como AAA-11A (sin espacios y en mayúsculas).', 
                            :if => "kind == 'Moto'"
    validates_length_of :plate, :minimum => 7, :maximum => 7, message: 'debe tener un formato correcto.'
    validates_uniqueness_of :plate, message: 'ya existe en el sistema.'
    validates_inclusion_of :kind, :in => ["Carro", "Moto"], message: 'de vehículo no es correcto.'
    validates :model, :length => { :minimum => 4, :maximum => 4, message: 'debe ser un número de 4 dígitos.' }
    validates_format_of :brand, with: /[a-zA-Z0-9 ]+/, message: 'debe ser válido (ej: Renault, Mazda).'
    validates_format_of :color, with: /[a-zA-Z ]+/, message: 'debe ser válido (ej: Rojo, Negro).'
    #-------------------------------------------------------------------------

    def self.myVehicles(user)
        Vehicle.where(:user_id => user).all
    end

    def self.findVehicle(plate)
        Vehicle.where(:plate => plate)
    end

    def self.countMyVehicles(user)
        Vehicle.where(:user_id => user).count
    end    
    
    def self.totalVehicles()
        Vehicle.count
    end


end
