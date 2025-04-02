export interface Props {
  name: string;
  description: string;
}

function IngredientCard({ name, description }: Props) {
  return (
    <div className="relative flex flex-col my-2 bg-white shadow-sm border border-slate-200 rounded-lg w-40">
      <div className="relative p-2 h-40 overflow-hidden rounded-xl bg-clip-border">
        <img
          data-tooltip-target="tooltip-description"
          src={`./imgs/ingredients/${name}.png`}
          alt="ingredient-image"
          className="h-full w-full object-cover rounded-md"
        />
      </div>
      <div
        id="tooltip-description"
        role="tooltip"
        className="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-xs opacity-0 tooltip dark:bg-gray-700"
      >
        {description}
        <div className="tooltip-arrow" data-popper-arrow></div>
      </div>
      <div className="p-4">
        <p className="text-slate-800 text-m font-semibold text-center">
          {name}
        </p>
      </div>
    </div>
  );
}

export default IngredientCard;
