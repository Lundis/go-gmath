package vec2

import "encoding/json"

/**
 * Custom JSON marshal/unmarshal to represent vec2 as [x, y] instead of {"X": x, "Y": y}
 */

func (v F) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]float32{v.X, v.Y})
}
func (v I) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]int32{v.X, v.Y})
}
func (v D) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]float64{v.X, v.Y})
}

func (v *F) UnmarshalJSON(data []byte) error {
	var tmp [2]float32
	if err := json.Unmarshal(data, &tmp); err != nil {
		type Alias F
		var alias Alias
		// fallback to default unmarshalling
		if err := json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*v = F(alias)
		return nil
	}
	v.X = tmp[0]
	v.Y = tmp[1]
	return nil
}

func (i *I) UnmarshalJSON(data []byte) error {
	var tmp [2]int32
	if err := json.Unmarshal(data, &tmp); err != nil {
		type Alias I
		var alias Alias
		// fallback to default unmarshalling
		if err := json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*i = I(alias)
		return nil
	}
	i.X = tmp[0]
	i.Y = tmp[1]
	return nil
}

func (v *D) UnmarshalJSON(data []byte) error {
	var tmp [2]float64
	if err := json.Unmarshal(data, &tmp); err != nil {
		type Alias D
		var alias Alias
		// fallback to default unmarshalling
		if err := json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*v = D(alias)
		return nil
	}
	v.X = tmp[0]
	v.Y = tmp[1]
	return nil
}
