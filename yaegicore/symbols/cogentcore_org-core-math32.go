// Code generated by 'yaegi extract cogentcore.org/core/math32'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/math32"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/math32/math32"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Abs":                      reflect.ValueOf(math32.Abs),
		"Acos":                     reflect.ValueOf(math32.Acos),
		"Acosh":                    reflect.ValueOf(math32.Acosh),
		"Asin":                     reflect.ValueOf(math32.Asin),
		"Asinh":                    reflect.ValueOf(math32.Asinh),
		"Atan":                     reflect.ValueOf(math32.Atan),
		"Atan2":                    reflect.ValueOf(math32.Atan2),
		"Atanh":                    reflect.ValueOf(math32.Atanh),
		"B2":                       reflect.ValueOf(math32.B2),
		"B2Empty":                  reflect.ValueOf(math32.B2Empty),
		"B2FromFixed":              reflect.ValueOf(math32.B2FromFixed),
		"B2FromRect":               reflect.ValueOf(math32.B2FromRect),
		"B3":                       reflect.ValueOf(math32.B3),
		"B3Empty":                  reflect.ValueOf(math32.B3Empty),
		"BarycoordFromPoint":       reflect.ValueOf(math32.BarycoordFromPoint),
		"Cbrt":                     reflect.ValueOf(math32.Cbrt),
		"Ceil":                     reflect.ValueOf(math32.Ceil),
		"ContainsPoint":            reflect.ValueOf(math32.ContainsPoint),
		"CopyFloat32s":             reflect.ValueOf(math32.CopyFloat32s),
		"CopyFloat64s":             reflect.ValueOf(math32.CopyFloat64s),
		"Copysign":                 reflect.ValueOf(math32.Copysign),
		"Cos":                      reflect.ValueOf(math32.Cos),
		"Cosh":                     reflect.ValueOf(math32.Cosh),
		"DegToRad":                 reflect.ValueOf(math32.DegToRad),
		"DegToRadFactor":           reflect.ValueOf(constant.MakeFromLiteral("0.0174532925199432957692369076848861271344287188854172545609719143893343406766598654219872641535175884721781352014070117566218413351995865581278454486637752166873317868068496601374750554214188014157116413116455078125", token.FLOAT, 0)),
		"Dim":                      reflect.ValueOf(math32.Dim),
		"DimsN":                    reflect.ValueOf(math32.DimsN),
		"DimsValues":               reflect.ValueOf(math32.DimsValues),
		"E":                        reflect.ValueOf(constant.MakeFromLiteral("2.71828182845904523536028747135266249775724709369995957496696762566337824315673231520670375558666729784504486779277967997696994772644702281675346915668215131895555530285035761295375777990557253360748291015625", token.FLOAT, 0)),
		"Erf":                      reflect.ValueOf(math32.Erf),
		"Erfc":                     reflect.ValueOf(math32.Erfc),
		"Erfcinv":                  reflect.ValueOf(math32.Erfcinv),
		"Erfinv":                   reflect.ValueOf(math32.Erfinv),
		"Exp":                      reflect.ValueOf(math32.Exp),
		"Exp2":                     reflect.ValueOf(math32.Exp2),
		"Expm1":                    reflect.ValueOf(math32.Expm1),
		"FMA":                      reflect.ValueOf(math32.FMA),
		"FastExp":                  reflect.ValueOf(math32.FastExp),
		"FitGeomInWindow":          reflect.ValueOf(math32.FitGeomInWindow),
		"Floor":                    reflect.ValueOf(math32.Floor),
		"Frexp":                    reflect.ValueOf(math32.Frexp),
		"FromFixed":                reflect.ValueOf(math32.FromFixed),
		"FromPoint":                reflect.ValueOf(math32.FromPoint),
		"Gamma":                    reflect.ValueOf(math32.Gamma),
		"Hypot":                    reflect.ValueOf(math32.Hypot),
		"Identity2":                reflect.ValueOf(math32.Identity2),
		"Identity3":                reflect.ValueOf(math32.Identity3),
		"Identity4":                reflect.ValueOf(math32.Identity4),
		"Ilogb":                    reflect.ValueOf(math32.Ilogb),
		"Inf":                      reflect.ValueOf(math32.Inf),
		"Infinity":                 reflect.ValueOf(&math32.Infinity).Elem(),
		"IntMultiple":              reflect.ValueOf(math32.IntMultiple),
		"IntMultipleGE":            reflect.ValueOf(math32.IntMultipleGE),
		"IsInf":                    reflect.ValueOf(math32.IsInf),
		"IsNaN":                    reflect.ValueOf(math32.IsNaN),
		"J0":                       reflect.ValueOf(math32.J0),
		"J1":                       reflect.ValueOf(math32.J1),
		"Jn":                       reflect.ValueOf(math32.Jn),
		"Ldexp":                    reflect.ValueOf(math32.Ldexp),
		"Lerp":                     reflect.ValueOf(math32.Lerp),
		"Lgamma":                   reflect.ValueOf(math32.Lgamma),
		"Ln10":                     reflect.ValueOf(constant.MakeFromLiteral("2.30258509299404568401799145468436420760110148862877297603332784146804725494827975466552490443295866962642372461496758838959542646932914211937012833592062802600362869664962772731087170541286468505859375", token.FLOAT, 0)),
		"Ln2":                      reflect.ValueOf(constant.MakeFromLiteral("0.6931471805599453094172321214581765680755001343602552541206800092715999496201383079363438206637927920954189307729314303884387720696314608777673678644642390655170150035209453154294578780536539852619171142578125", token.FLOAT, 0)),
		"Log":                      reflect.ValueOf(math32.Log),
		"Log10":                    reflect.ValueOf(math32.Log10),
		"Log10E":                   reflect.ValueOf(constant.MakeFromLiteral("0.43429448190325182765112891891660508229439700580366656611445378416636798190620320263064286300825210972160277489744884502676719847561509639618196799746596688688378591625127711495224502868950366973876953125", token.FLOAT, 0)),
		"Log1p":                    reflect.ValueOf(math32.Log1p),
		"Log2":                     reflect.ValueOf(math32.Log2),
		"Log2E":                    reflect.ValueOf(constant.MakeFromLiteral("1.44269504088896340735992468100189213742664595415298593413544940772066427768997545329060870636212628972710992130324953463427359402479619301286929040235571747101382214539290471666532766903401352465152740478515625", token.FLOAT, 0)),
		"Logb":                     reflect.ValueOf(math32.Logb),
		"Matrix3FromMatrix2":       reflect.ValueOf(math32.Matrix3FromMatrix2),
		"Matrix3FromMatrix4":       reflect.ValueOf(math32.Matrix3FromMatrix4),
		"Matrix3Rotate2D":          reflect.ValueOf(math32.Matrix3Rotate2D),
		"Matrix3Scale2D":           reflect.ValueOf(math32.Matrix3Scale2D),
		"Matrix3Translate2D":       reflect.ValueOf(math32.Matrix3Translate2D),
		"Max":                      reflect.ValueOf(math32.Max),
		"MaxFloat32":               reflect.ValueOf(constant.MakeFromLiteral("340282346638528859811704183484516925440", token.FLOAT, 0)),
		"MaxPos":                   reflect.ValueOf(math32.MaxPos),
		"Min":                      reflect.ValueOf(math32.Min),
		"MinPos":                   reflect.ValueOf(math32.MinPos),
		"Mod":                      reflect.ValueOf(math32.Mod),
		"Modf":                     reflect.ValueOf(math32.Modf),
		"NaN":                      reflect.ValueOf(math32.NaN),
		"NewArrayF32":              reflect.ValueOf(math32.NewArrayF32),
		"NewArrayU32":              reflect.ValueOf(math32.NewArrayU32),
		"NewEulerAnglesFromMatrix": reflect.ValueOf(math32.NewEulerAnglesFromMatrix),
		"NewFrustum":               reflect.ValueOf(math32.NewFrustum),
		"NewFrustumFromMatrix":     reflect.ValueOf(math32.NewFrustumFromMatrix),
		"NewLine2":                 reflect.ValueOf(math32.NewLine2),
		"NewLine3":                 reflect.ValueOf(math32.NewLine3),
		"NewLookAt":                reflect.ValueOf(math32.NewLookAt),
		"NewPlane":                 reflect.ValueOf(math32.NewPlane),
		"NewQuat":                  reflect.ValueOf(math32.NewQuat),
		"NewQuatAxisAngle":         reflect.ValueOf(math32.NewQuatAxisAngle),
		"NewQuatEuler":             reflect.ValueOf(math32.NewQuatEuler),
		"NewRay":                   reflect.ValueOf(math32.NewRay),
		"NewSphere":                reflect.ValueOf(math32.NewSphere),
		"NewTriangle":              reflect.ValueOf(math32.NewTriangle),
		"NewVector3Color":          reflect.ValueOf(math32.NewVector3Color),
		"NewVector4Color":          reflect.ValueOf(math32.NewVector4Color),
		"Nextafter":                reflect.ValueOf(math32.Nextafter),
		"Normal":                   reflect.ValueOf(math32.Normal),
		"OtherDim":                 reflect.ValueOf(math32.OtherDim),
		"ParseAngle32":             reflect.ValueOf(math32.ParseAngle32),
		"ParseFloat32":             reflect.ValueOf(math32.ParseFloat32),
		"Phi":                      reflect.ValueOf(constant.MakeFromLiteral("1.6180339887498948482045868343656381177203091798057628621354486119746080982153796619881086049305501566952211682590824739205931370737029882996587050475921915678674035433959321750307935872115194797515869140625", token.FLOAT, 0)),
		"Pi":                       reflect.ValueOf(constant.MakeFromLiteral("3.141592653589793238462643383279502884197169399375105820974944594789982923695635954704435713335896673485663389728754819466702315787113662862838515639906529162340867271374644786874341662041842937469482421875", token.FLOAT, 0)),
		"PointDim":                 reflect.ValueOf(math32.PointDim),
		"PointsCheckN":             reflect.ValueOf(math32.PointsCheckN),
		"Pow":                      reflect.ValueOf(math32.Pow),
		"Pow10":                    reflect.ValueOf(math32.Pow10),
		"RadToDeg":                 reflect.ValueOf(math32.RadToDeg),
		"RadToDegFactor":           reflect.ValueOf(constant.MakeFromLiteral("57.295779513082320876798154814105170332405472466564321549160243902428585054360559672397261399470815487380868161395148776362013889310162423528726959840779630006155203887467652901221981665003113448619842529296875", token.FLOAT, 0)),
		"ReadPoints":               reflect.ValueOf(math32.ReadPoints),
		"RectFromPosSizeMax":       reflect.ValueOf(math32.RectFromPosSizeMax),
		"RectFromPosSizeMin":       reflect.ValueOf(math32.RectFromPosSizeMin),
		"RectInNotEmpty":           reflect.ValueOf(math32.RectInNotEmpty),
		"Remainder":                reflect.ValueOf(math32.Remainder),
		"Rotate2D":                 reflect.ValueOf(math32.Rotate2D),
		"Round":                    reflect.ValueOf(math32.Round),
		"RoundToEven":              reflect.ValueOf(math32.RoundToEven),
		"SRGBFromLinear":           reflect.ValueOf(math32.SRGBFromLinear),
		"SRGBToLinear":             reflect.ValueOf(math32.SRGBToLinear),
		"Scale2D":                  reflect.ValueOf(math32.Scale2D),
		"SetPointDim":              reflect.ValueOf(math32.SetPointDim),
		"Shear2D":                  reflect.ValueOf(math32.Shear2D),
		"Sign":                     reflect.ValueOf(math32.Sign),
		"Signbit":                  reflect.ValueOf(math32.Signbit),
		"Sin":                      reflect.ValueOf(math32.Sin),
		"Sincos":                   reflect.ValueOf(math32.Sincos),
		"Sinh":                     reflect.ValueOf(math32.Sinh),
		"Skew2D":                   reflect.ValueOf(math32.Skew2D),
		"SmallestNonzeroFloat32":   reflect.ValueOf(constant.MakeFromLiteral("1.40129846432481707092372958328991613128026194187651577175706828388979108268586060148663818836212158203125e-45", token.FLOAT, 0)),
		"Sqrt":                     reflect.ValueOf(math32.Sqrt),
		"Sqrt2":                    reflect.ValueOf(constant.MakeFromLiteral("1.414213562373095048801688724209698078569671875376948073176679739576083351575381440094441524123797447886801949755143139115339040409162552642832693297721230919563348109313505318596071447245776653289794921875", token.FLOAT, 0)),
		"SqrtE":                    reflect.ValueOf(constant.MakeFromLiteral("1.64872127070012814684865078781416357165377610071014801157507931167328763229187870850146925823776361770041160388013884200789716007979526823569827080974091691342077871211546646890155898290686309337615966796875", token.FLOAT, 0)),
		"SqrtPhi":                  reflect.ValueOf(constant.MakeFromLiteral("1.2720196495140689642524224617374914917156080418400962486166403754616080542166459302584536396369727769747312116100875915825863540562126478288118732191412003988041797518382391984914647764526307582855224609375", token.FLOAT, 0)),
		"SqrtPi":                   reflect.ValueOf(constant.MakeFromLiteral("1.772453850905516027298167483341145182797549456122387128213807789740599698370237052541269446184448945647349951047154197675245574635259260134350885938555625028620527962319730619356050738133490085601806640625", token.FLOAT, 0)),
		"Tan":                      reflect.ValueOf(math32.Tan),
		"Tanh":                     reflect.ValueOf(math32.Tanh),
		"ToFixed":                  reflect.ValueOf(math32.ToFixed),
		"ToFixedPoint":             reflect.ValueOf(math32.ToFixedPoint),
		"Translate2D":              reflect.ValueOf(math32.Translate2D),
		"Trunc":                    reflect.ValueOf(math32.Trunc),
		"Truncate":                 reflect.ValueOf(math32.Truncate),
		"Truncate64":               reflect.ValueOf(math32.Truncate64),
		"Vec2":                     reflect.ValueOf(math32.Vec2),
		"Vec2i":                    reflect.ValueOf(math32.Vec2i),
		"Vec3":                     reflect.ValueOf(math32.Vec3),
		"Vec3i":                    reflect.ValueOf(math32.Vec3i),
		"Vec4":                     reflect.ValueOf(math32.Vec4),
		"Vector2FromFixed":         reflect.ValueOf(math32.Vector2FromFixed),
		"Vector2Scalar":            reflect.ValueOf(math32.Vector2Scalar),
		"Vector2iScalar":           reflect.ValueOf(math32.Vector2iScalar),
		"Vector3FromVector4":       reflect.ValueOf(math32.Vector3FromVector4),
		"Vector3Scalar":            reflect.ValueOf(math32.Vector3Scalar),
		"Vector3iScalar":           reflect.ValueOf(math32.Vector3iScalar),
		"Vector4FromVector3":       reflect.ValueOf(math32.Vector4FromVector3),
		"Vector4Scalar":            reflect.ValueOf(math32.Vector4Scalar),
		"W":                        reflect.ValueOf(math32.W),
		"X":                        reflect.ValueOf(math32.X),
		"Y":                        reflect.ValueOf(math32.Y),
		"Y0":                       reflect.ValueOf(math32.Y0),
		"Y1":                       reflect.ValueOf(math32.Y1),
		"Yn":                       reflect.ValueOf(math32.Yn),
		"Z":                        reflect.ValueOf(math32.Z),

		// type definitions
		"ArrayF32":  reflect.ValueOf((*math32.ArrayF32)(nil)),
		"ArrayU32":  reflect.ValueOf((*math32.ArrayU32)(nil)),
		"Box2":      reflect.ValueOf((*math32.Box2)(nil)),
		"Box3":      reflect.ValueOf((*math32.Box3)(nil)),
		"Dims":      reflect.ValueOf((*math32.Dims)(nil)),
		"Frustum":   reflect.ValueOf((*math32.Frustum)(nil)),
		"Geom2DInt": reflect.ValueOf((*math32.Geom2DInt)(nil)),
		"Line2":     reflect.ValueOf((*math32.Line2)(nil)),
		"Line3":     reflect.ValueOf((*math32.Line3)(nil)),
		"Matrix2":   reflect.ValueOf((*math32.Matrix2)(nil)),
		"Matrix3":   reflect.ValueOf((*math32.Matrix3)(nil)),
		"Matrix4":   reflect.ValueOf((*math32.Matrix4)(nil)),
		"Plane":     reflect.ValueOf((*math32.Plane)(nil)),
		"Quat":      reflect.ValueOf((*math32.Quat)(nil)),
		"Ray":       reflect.ValueOf((*math32.Ray)(nil)),
		"Sphere":    reflect.ValueOf((*math32.Sphere)(nil)),
		"Triangle":  reflect.ValueOf((*math32.Triangle)(nil)),
		"Vector2":   reflect.ValueOf((*math32.Vector2)(nil)),
		"Vector2i":  reflect.ValueOf((*math32.Vector2i)(nil)),
		"Vector3":   reflect.ValueOf((*math32.Vector3)(nil)),
		"Vector3i":  reflect.ValueOf((*math32.Vector3i)(nil)),
		"Vector4":   reflect.ValueOf((*math32.Vector4)(nil)),
	}
}
