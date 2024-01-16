import './App.css'
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue
} from "@/components/ui/select.tsx";



function App() {

  return (
    <>
        <Select>
            <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Choissisez un aéroport" />
            </SelectTrigger>
            <SelectContent>
                <SelectGroup>
                    <SelectLabel>Aéroports</SelectLabel>
                    <SelectItem value="NTE">Nantes</SelectItem>
                    <SelectItem value="ORY">Paris Orly</SelectItem>
                    <SelectItem value="CDG">Paris Charles de Gaulle</SelectItem>
                    <SelectItem value="DBX">Dubai</SelectItem>
                </SelectGroup>
            </SelectContent>
        </Select>
    </>
  )
}

export default App
