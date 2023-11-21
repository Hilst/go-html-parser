
def read_env_file(file_name: str) -> dict[str, str]:
   env_dict = {}
   with open(file_name, 'r') as file:
       for line in file:
           # Ignore comments and empty lines
           if not line.startswith('#') and line.strip():
               key, value = line.strip().split('=', 1)
               env_dict[key] = value
   return env_dict